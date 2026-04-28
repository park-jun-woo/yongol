import * as React from 'react'
export const Input = React.forwardRef<HTMLInputElement, any>((props, ref) => <input ref={ref} {...props} />)
