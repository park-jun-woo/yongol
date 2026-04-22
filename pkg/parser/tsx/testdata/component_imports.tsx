import React from 'react'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { CreditsGauge } from './components/CreditsGauge'
import clsx from 'clsx'

export default function SettingsPage() {
  return (
    <Card>
      <CreditsGauge summary={{ credits_balance: 10, plan_type: 'free' }} />
      <Button className={clsx('primary')}>Save</Button>
    </Card>
  )
}
