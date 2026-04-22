#!/usr/bin/env python3
"""
Dummy SMTP server for yongol hurl tests.
Accepts all emails and discards them. Supports AUTH PLAIN for Go smtp.PlainAuth.

Usage:
    python3 dummy-smtp.py [port]
    default port: 2525

Environment variables for the backend server:
    SMTP_HOST=127.0.0.1
    SMTP_PORT=2525
    SMTP_USERNAME=test
    SMTP_PASSWORD=test
    SMTP_FROM=test@test.com
"""

import socket
import sys
import threading


def handle(conn, addr):
    conn.sendall(b"220 localhost SMTP\r\n")
    while True:
        try:
            data = conn.recv(4096)
        except ConnectionResetError:
            break
        if not data:
            break
        lines = data.decode(errors="replace").strip().split("\r\n")
        for line in lines:
            cmd = line.strip().upper()
            if cmd.startswith("EHLO") or cmd.startswith("HELO"):
                conn.sendall(b"250-localhost\r\n250-AUTH PLAIN LOGIN\r\n250 OK\r\n")
            elif cmd.startswith("AUTH"):
                conn.sendall(b"235 Authentication successful\r\n")
            elif cmd.startswith("MAIL"):
                conn.sendall(b"250 OK\r\n")
            elif cmd.startswith("RCPT"):
                conn.sendall(b"250 OK\r\n")
            elif cmd.startswith("DATA"):
                conn.sendall(b"354 Go\r\n")
                buf = b""
                while True:
                    d = conn.recv(4096)
                    buf += d
                    if buf.endswith(b"\r\n.\r\n"):
                        break
                conn.sendall(b"250 OK\r\n")
            elif cmd.startswith("QUIT"):
                conn.sendall(b"221 Bye\r\n")
                conn.close()
                return
            elif cmd:
                conn.sendall(b"250 OK\r\n")
    conn.close()


def main():
    port = int(sys.argv[1]) if len(sys.argv) > 1 else 2525
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    s.bind(("127.0.0.1", port))
    s.listen(5)
    print(f"Dummy SMTP listening on 127.0.0.1:{port}")
    sys.stdout.flush()
    while True:
        conn, addr = s.accept()
        threading.Thread(target=handle, args=(conn, addr), daemon=True).start()


if __name__ == "__main__":
    main()
