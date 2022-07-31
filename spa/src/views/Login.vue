<template>
  <div class="is-flex is-flex-direction-column is-full-vh is-full-vw">
    <div class="is-flex-grow-1">&nbsp;</div>
    <div class="login">
      <form class="card is-shadowless py-5 px-5" @submit.prevent="login">
        <form-field for="email" label="Login" icon="fad-user" @changed="checkValidation" :validation="validation">
          <input class="input" v-model="userLogin.email" name="email" type="text" autocomplete="none" @blur="loadSites" />
        </form-field>
        <form-field for="password" label="Password" icon="fad-lock" class="mt-4" @changed="checkValidation" :validation="validation">
          <input class="input" v-model="userLogin.password" name="password" type="password" />
        </form-field>
        <form-field for="siteUlid" label="Site" icon="fad-home-lg-alt" class="mt-4" @changed="checkValidation" :validation="validation">
          <drp-down name="siteUlid" v-model="userLogin.siteUlid" :data="sites" label-key="name" value-key="siteUlid" :disabled="sites.length < 1" />
        </form-field>
        <!-- <loading class="text-blue-800 bg-grey-400 bg-opacity-75" :on="['Login']" /> -->
        <o-button @click="login" class="mt-4 is-blue float-right">Login</o-button>
      </form>
    </div>
    <div class="is-flex-grow-1">&nbsp;</div>
  </div>
</template>

<script lang="ts" setup>
import $api from '@/api'
import { Site, UnauthorisedUser } from '@/api/pb/access.pb'
import { useUserAccessStore } from '@/store/userAccess'
import validate, { Validation } from '@/utils/validate'
import { Ref, ref } from 'vue'

const userAccess = useUserAccessStore()

let sites: Ref<Array<Site>> = ref([])
let userLogin: Ref<UnauthorisedUser> = ref({
  email: 'josh@nerdy.co.nz',
  password: 'yellowst4r',
  siteUlid: '01EHZXH0YBCM8Q8PEFDZB8K3WW',
})

async function login() {
  await checkValidation()
  if (validation.value.isValid) {
    await userAccess.login(userLogin.value)
  }
}

async function loadSites() {
  // let result = await validate(this.userLogin, null, {
  //     email: { presence: true, email: true }
  //   },
  //   'email'
  // )
  // if (result.isValid) {
  let resp = await $api.access.ValidSites({
    email: userLogin.value.email,
  })
  sites.value = resp.sites
  if (sites.value.length === 1) {
    userLogin.value.siteUlid = sites.value[0].siteUlid
  }
  // }
}

let validation: Ref<Validation> = ref({ isValid: false })
async function checkValidation(fieldName = '') {
  let result = await validate(
    userLogin.value,
    validation.value,
    {
      email: { presence: true, email: true },
      password: { presence: { allowEmpty: false } },
      siteUlid: { presence: { allowEmpty: false, message: '^ choose which site to login to' } },
    },
    fieldName
  )
  validation.value = result
}
</script>
<style lang="scss">
  .login {
    width: 640px;
    margin: 0 auto;
  }
</style>