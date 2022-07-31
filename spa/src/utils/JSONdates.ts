export default {
  dtrx2: /\d{4}-\d{2}-\d{2}/,
  parse(obj: string | null) {
    if (!obj) {
      return null
    }
    var parsedObj = JSON.parse(obj)
    return this.parseDates(parsedObj)
  },
  parseDates(obj: any) {
    // iterate properties
    for (let pName in obj) {
      let value = obj[pName]
      // make sure the property is 'truthy'
      if (value) {
        if (typeof value === 'string' && this.dtrx2.test(value)) {
          // parse and replace
          obj[pName] = new Date(obj[pName])
        }
        // determine if the property is an array
        else if (Array.isArray(value)) {
          for (let i = 0; i < value.length; i++) {
            this.parseDates(value[i])
          }
        }
        // determine if the property is an object
        else if (typeof value == 'object') {
          this.parseDates(value)
        }
      }
    }
    return obj
  },
}
