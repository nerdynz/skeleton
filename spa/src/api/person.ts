import { useUserAccessStore } from '@/store/userAccess'
import validate, { Validation } from '@/utils/validate'
import { ulid } from 'ulid'
import { Person, DeletePerson, LoadPerson, PagedPeople, SavePerson } from './pb/person.pb'

const defaultPersonValidation = {
  siteUlid: { presence: { allowEmpty: false } },
  personUlid: { presence: { allowEmpty: false } },
}

const personApi = {
  createPerson: () => {
    return {
      personUlid: ulid(),
      siteUlid: useUserAccessStore().siteUlid,
    } as Person
  },
  savePerson: SavePerson,
  pagedPeople: PagedPeople,
  loadPerson: LoadPerson,
  deletePerson: DeletePerson,
  validate: (record: Person | null, validation: Validation | null, fieldName: string = '', validationDefinition?: any) => {
    validationDefinition = validationDefinition || defaultPersonValidation
    return validate(record, validation, validationDefinition, fieldName)
  },
}
export default personApi