export const isEmptyObject = (object: any): boolean => {
    for (const key in object) return false
    return true
  }