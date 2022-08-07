import * as validatejs from "validate.js"

export interface Validation {
  isValid: boolean,
  result?: Record<string, Array<string>>
}

export default function validate(record: any | null, validation: Validation | null, validationDefinition: any, fieldName: string = '') : Promise<Validation> {
  let result: Record<string, any> = {}
  if (fieldName) {
    // we are checking an indiviual field so ensure the previous result gets passed
    result = validation?.result || {}
    console.log('result', result)
    if (result.hasOwnProperty(fieldName)) {
      delete result[fieldName]
    }
  }
  return new Promise(async (resolve, reject) => {
    if (!record) {
      resolve({
        isValid: false,
      })
    }
    try {
      let toValidate: any = validationDefinition
      if (fieldName in validationDefinition) {
        toValidate = { [fieldName]: validationDefinition[fieldName] }
        if ('equality' in toValidate[fieldName]) {
          let equalityFieldName = typeof(toValidate[fieldName].equality) === 'string' ? toValidate[fieldName].equality : toValidate[fieldName].equality.attribute
          toValidate[equalityFieldName] = validationDefinition[equalityFieldName] 
        }
      }
      let newRes = await validatejs.async(record, toValidate)
      let validationResult = {
        isValid: Object.keys(result).length === 0,
        result: result
      }
      resolve(validationResult)
    } catch (validationErrors: any) {
      console.warn('Validation Error', validationErrors)
      if (fieldName) {
        result[fieldName] = validationErrors[fieldName]
      } else {
        result = validationErrors as any
      }
      resolve({
        isValid: false,
        result: result
      })
    }
  })
}