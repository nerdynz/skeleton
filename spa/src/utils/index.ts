import { isEmpty as _isEmpty } from 'validate.js'

export function debounce(fn: (...params: any[]) => any, n: number = 250, immed: boolean = false) {
  let timer: number | undefined = undefined
  return function (this: any, ...args: any[]) {
    if (timer === undefined && immed) {
      fn.apply(this, args)
    }
    clearTimeout(timer)
    timer = setTimeout(() => fn.apply(this, args), n)
    return timer
  }
}

export function isEmpty(value: any): boolean {
  return _isEmpty(value)
}

export function isPagedData(pagedData?: any) {
  if (pagedData) {
    return 'total' in pagedData && 'direction' in pagedData && 'sort' in pagedData
  }
  return false
}