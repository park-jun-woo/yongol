import * as React from 'react'
import { cn } from '@/lib/utils'

export type ModalProps = {
  open: boolean
  onClose: () => void
  children?: React.ReactNode
  className?: string
}

export function Modal({ open, onClose, className, children }: ModalProps) {
  if (!open) return null
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-foreground/50" onClick={onClose}>
      <div className={cn('bg-background text-foreground rounded-lg border border-border p-6 shadow-lg min-w-[320px]', className)} onClick={e => e.stopPropagation()}>
        {children}
      </div>
    </div>
  )
}
