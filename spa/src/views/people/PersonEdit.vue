<template>
  <slideout :title="pageTitle" ref="popover" :actions="actions" @closed="$router.push({ name: 'PersonList' })">
    <form v-if="record" @keyup.enter="save">
      <form-field class="mt-4" :validation="validation" @changed="validate" for="name" label="Name">
        <o-input label="Name" name="name" type="text" v-model="record.name" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="username" label="Username">
        <o-input label="Username" name="username" type="text" v-model="record.username" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="email" label="Email">
        <o-input label="Email" name="email" type="email" v-model="record.email" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="phone" label="Phone">
        <o-input label="Phone" name="phone" type="text" v-model="record.phone" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="role" label="Role">
        <o-input label="Role" name="role" type="text" v-model="record.role" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="initials" label="Initials" wraps-richtext>
        <rich-text label="Initials" name="initials" type="text" v-model="record.initials" />
      </form-field>
      <form-field class="mt-4" :validation="validation" @changed="validate" for="password" label="Password">
        <o-input label="Password" name="password" type="text" v-model="record.password" />
      </form-field>
    </form>
    <loading :on="['LoadPerson']" />
  </slideout>
</template>

<script lang="ts" setup>
import api from '@/api'
import { Person } from '@/api/pb/person.pb'
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
  return `${isNew.value ? 'Create' : 'Edit'} Person`
})

async function validate(fieldName = '') {
  validation.value = await api.person.validate(record.value!, validation.value, fieldName)
}

const isValid = computed(() => {
  return validation.value.isValid
})

let popover = ref<InstanceType<typeof Slideout> | null>(null)
let record: Ref<Person | null> = ref(null)

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

onMounted(() => {
  load()
})
</script>
