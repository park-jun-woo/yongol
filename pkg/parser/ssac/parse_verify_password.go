//ff:func feature=ssac-parse type=parser control=sequence topic=timing-defense
//ff:what @verify-password — 로그인 타이밍 방어 묶음 시퀀스 파싱

package ssac

import (
	"fmt"
	"strconv"
	"strings"
)

// parseVerifyPassword parses the @verify-password syntax:
//
//	@verify-password <Model>.<emailCol>=<emailExpr> <Model>.<hashCol> vs <pwExpr>
//	  -> <var> <status> "<message>"
//
// Example:
//
//	@verify-password User.email=request.body.email User.password_hash vs request.body.password
//	  -> user 401 "Invalid credentials"
//
// Semantics (handled by the code generator):
//   1. findByEmail against Model using emailCol/emailExpr
//   2. on miss: bcrypt-compare pwExpr against auth.DummyHash (timing equaliser) → status+message
//   3. on hit: bcrypt-compare pwExpr against row.<hashCol> → status+message on mismatch
//   4. on success: bind the user row to <var>
//
// The Model must be a sqlc row type (e.g. User) and both emailCol/hashCol
// must be actual columns on that table — validation layer enforces this.
func parseVerifyPassword(rest string) (*Sequence, error) {
	rest = strings.TrimSpace(rest)
	seq := &Sequence{Type: SeqVerifyPassword}

	// Split off the trailing "-> <var> <status> <msg>" clause.
	arrowIdx := strings.Index(rest, "->")
	if arrowIdx < 0 {
		return nil, fmt.Errorf("@verify-password: missing '->' clause")
	}
	lhs := strings.TrimSpace(rest[:arrowIdx])
	rhs := strings.TrimSpace(rest[arrowIdx+2:])

	// Parse LHS: <Model>.<emailCol>=<emailExpr>  <Model>.<hashCol> vs <pwExpr>
	// Split around " vs " first — pwExpr sits after.
	vsIdx := strings.Index(lhs, " vs ")
	if vsIdx < 0 {
		return nil, fmt.Errorf("@verify-password: missing 'vs' separator")
	}
	before := strings.TrimSpace(lhs[:vsIdx])
	seq.PasswordExpr = strings.TrimSpace(lhs[vsIdx+4:])
	if seq.PasswordExpr == "" {
		return nil, fmt.Errorf("@verify-password: password expression empty")
	}

	// `before` now holds: <Model>.<emailCol>=<emailExpr>  <Model>.<hashCol>
	// Split on whitespace — first whitespace separates email clause from hash clause.
	beforeParts := strings.Fields(before)
	if len(beforeParts) < 2 {
		return nil, fmt.Errorf("@verify-password: email/hash clauses missing")
	}
	emailClause := beforeParts[0]
	hashClause := strings.Join(beforeParts[1:], " ")

	// Email clause: <Model>.<emailCol>=<emailExpr>
	eqIdx := strings.Index(emailClause, "=")
	if eqIdx < 0 {
		return nil, fmt.Errorf("@verify-password: email clause missing '='")
	}
	emailKey := strings.TrimSpace(emailClause[:eqIdx])
	seq.EmailExpr = strings.TrimSpace(emailClause[eqIdx+1:])
	emailKeyParts := strings.SplitN(emailKey, ".", 2)
	if len(emailKeyParts) != 2 {
		return nil, fmt.Errorf("@verify-password: email clause key must be Model.column")
	}
	seq.Model = strings.TrimSpace(emailKeyParts[0])
	seq.EmailCol = strings.TrimSpace(emailKeyParts[1])

	// Hash clause: <Model>.<hashCol> — Model must match the email clause.
	hashParts := strings.SplitN(hashClause, ".", 2)
	if len(hashParts) != 2 {
		return nil, fmt.Errorf("@verify-password: hash clause must be Model.column")
	}
	if strings.TrimSpace(hashParts[0]) != seq.Model {
		return nil, fmt.Errorf("@verify-password: email Model %q and hash Model %q must match", seq.Model, hashParts[0])
	}
	seq.HashCol = strings.TrimSpace(hashParts[1])

	// Parse RHS: <var> <status> "<msg>"
	// Find the first quote — everything before it is "<var> <status>".
	preMsg := rhs
	if q := strings.Index(rhs, `"`); q >= 0 {
		preMsg = strings.TrimSpace(rhs[:q])
		afterQuote := rhs[q:]
		msg, _ := extractQuoted(afterQuote)
		seq.Message = msg
	}

	rhsTokens := strings.Fields(preMsg)
	if len(rhsTokens) < 2 {
		return nil, fmt.Errorf("@verify-password: RHS must be '<var> <status> \"<msg>\"'")
	}
	varName := rhsTokens[0]
	seq.Result = &Result{Var: varName, Type: seq.Model}
	if code, err := strconv.Atoi(rhsTokens[1]); err == nil && code > 0 {
		seq.ErrStatus = code
	} else {
		return nil, fmt.Errorf("@verify-password: status must be integer, got %q", rhsTokens[1])
	}

	return seq, nil
}
