<template>
  <slideout :title="pageTitle" ref="popover" :actions="actions" @closed="$router.push({ name: 'PersonList' })">
    <form v-if="record" @keyup.enter="save">
      <form-field class="mt-4" :validation="validation" @changed="validate" for="name" label="Name">
        <o-input name="name" type="text" v-model="record.name" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="username" label="Username">
        <o-input name="username" type="text" v-model="record.username" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="email" label="Email">
        <o-input name="email" type="email" v-model="record.email" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="phone" label="Phone">
        <o-input name="phone" type="text" v-model="record.phone" />
      </form-field>
      <!-- <form-field class="mt-4" :validation="validation" @changed="validate" for="role" label="Role">
        <o-input label="Role" name="role" type="text" v-model="record.role" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="initials" label="Initials" wraps-richtext>
        <rich-text label="Initials" name="initials" type="text" v-model="record.initials" />
      </form-field> -->
      <section class="password-change">
        <form-field class="mt-4" :validation="validation" @changed="validate" for="password" label="Password">
          <o-input type="password" name="password" v-model="record.password" :disabled="!isChangingPassword" />
          <button v-if="!isChangingPassword" type="button" class="button is-small password-change-btn" @click="changePassword">Change</button>
        </form-field>
        <form-field v-if="isChangingPassword" class="mt-4" :validation="validation" @changed="validate" for="passwordConfirmation" label="Confirm Password">
          <o-input type="password" name="passwordConfirmation" v-model="passwordChangeConfirmation" />
        </form-field>
      </section>
    </form>
    <loading :on="['LoadPerson']" />
  </slideout>
</template>

<script lang="ts" setup>
import api from '@/api'
import { Person } from '@/api/pb/person.pb'
import { PersonPasswordConfirmation } from '@/api/person'
import { Validation } from '@/utils/validate'
import Slideout from '@nerdynz/componenty/src/components/Popover/Slideout.vue'
import { computed, onMounted, ref, Ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const isNew = computed(() => {
  return !route.params?.ulid
})

const pageTitle = computed(() => {
  return `${isNew.value ? 'Create' : 'Edit'} User`
})

async function validate(fieldName = '') {
  let toValidate: Person | PersonPasswordConfirmation = record.value!
  let validateRules: any = {}
  if (isChangingPassword.value) {
    toValidate = {
      ...record.value!,
      passwordConfirmation: passwordChangeConfirmation.value,
    }
    validateRules.passwordConfirmation = {
      equality: 'password'
    }
  }
  validation.value = await api.person.validate(toValidate, validation.value, fieldName, validateRules)
}

const isValid = computed(() => {
  return validation.value.isValid
})

let popover = ref<InstanceType<typeof Slideout> | null>(null)
let record: Ref<Person | null> = ref(null)
let isChangingPassword = ref(false)
let passwordChangeConfirmation = ref('')

async function load() {
  record.value = null
  if (route.params?.ulid) {
    // EDIT
    record.value = await api.person.loadPerson({
      ulid: route.params?.ulid as string,
    })
  } else {
    record.value = api.person.createPerson()
  }
}

let validation: Ref<Validation> = ref({
  isValid: false,
})

const actions = computed(() => {
  return [
    {
      name: isNew.value ? 'Create' : 'Save',
      action: async () => {
        let saved = await save()
        if (saved) {
          popover.value!.close()
        }
      },
    },
    {
      name: 'Cancel',
      class: 'is-light',
      action: () => {
        popover.value!.close()
      },
    },
    {
      name: 'Delete',
      class: 'is-danger is-right',
      action: () => {
        api.person.deletePerson({
          ulid: route?.params?.ulid as string,
        })
        popover.value!.close()
      },
    },
  ]
})

async function save(): Promise<boolean> {
  await validate()
  if (isValid.value) {
    let person = await api.person.savePerson(record.value!)
    record.value = person
    router.replace({ name: 'PersonEdit', params: { ulid: person.personUlid } })
    return true
  }
  return false
}

function changePassword() {
  isChangingPassword.value = true
  record.value!.password = ''
}

onMounted(() => {
  load()
  if (isNew.value) {
    isChangingPassword.value = true
  }
})
</script>
<style lang="scss">
.password-change {
  position: relative;
  .password-change-btn {
    position: absolute;
    top: 0.5rem;
    right: 0.5rem;
  }
}
</style>
