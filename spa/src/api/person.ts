import { useUserAccessStore } from '@/store/userAccess'
import validate, { Validation } from '@/utils/validate'
import { ulid } from 'ulid'
import { DeletePerson, LoadPerson, PagedPeople, Person, SavePerson } from './pb/person.pb'

export interface PersonPasswordConfirmation extends Person {
  passwordConfirmation: string
}

const defaultPersonValidation = {
  siteUlid: { presence: { allowEmpty: false } },
  personUlid: { presence: { allowEmpty: false } },
  name: {
    presence: { allowEmpty: false },
    length: {
      minimum: 4,
    },
  },
  username: {
    presence: { allowEmpty: false },
    length: {
      minimum: 4,
    },
  },
  email: { email: true, presence: { allowEmpty: false } },
  password: {
    length: {
      minimum: 8,
    },
    presence: { allowEmpty: false },
  },
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
  validate: (record: Person | PersonPasswordConfirmation | null, validation: Validation | null, fieldName: string = '', validationDefinition?: any) => {
    let def = defaultPersonValidation
    if (validationDefinition) {
      def = {
        ...defaultPersonValidation,
        ...validationDefinition,
      }
    }
    return validate(record, validation, def, fieldName)
  },
}
export default personApi
