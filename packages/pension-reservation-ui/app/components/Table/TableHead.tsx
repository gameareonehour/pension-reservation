import { FC, ReactNode } from 'react'

const TableHead: FC<{ children: ReactNode }> = ({ children }) => {
  return <thead>{children}</thead>
}

export default TableHead
