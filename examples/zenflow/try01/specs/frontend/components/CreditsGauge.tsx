import React from 'react'
export type CreditsGaugeProps = { summary: { credits_balance: number; plan_type: string } }
export function CreditsGauge({ summary }: CreditsGaugeProps) {
  return <div>Credits: {summary.credits_balance} ({summary.plan_type})</div>
}
