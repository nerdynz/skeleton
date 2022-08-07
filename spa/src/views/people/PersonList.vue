<template>
  <with-actions class="fh-card p-4" :actions="actions" :paged-info="pagedInfo" @pagination-changed="changePageInfo">
    <template #left>
      <h1 class="title is-3 mb-0">Users</h1>
    </template>
    <template #right>
      <form-field for="search" icon="free-search" label="Search" class="with-button">
        <input class="input" name="search" v-model.trim="pagedInfo.search" type="text" @keyup.enter="loadData" @keyup.esc="pagedInfo.search = ''" autocomplete="off" />
      </form-field>
    </template>
    <data-table class="is-fullwidth is-stripped" :columns="columns" :rows="rows" :paged-info="pagedInfo" hide-pagination>
      <template #colgroup>
        <colgroup>
          <col width="5%" />
          <col width="15%" />
          <col width="15%" />
          <col width="20%" />
          <col width="20%" />
          <col width="12.5%" />
          <col width="12.5%" />
        </colgroup>
      </template>
      <template #actions="{ row }">
        <o-button tag="a" @click="edit(row.personUlid)" class="is-light" size="small" icon-pack="free" icon-left="edit is-small">Edit</o-button>
      </template>
      <loading :on="['PagedPeople']" />
    </data-table>
  </with-actions>
  <router-view />
</template>

<script lang="ts" setup>
import api from '@/api'
import { PagedInfo } from '@/api/pb/base.pb'
import { Person } from '@/api/pb/person.pb'
import ColumnMeta from '@nerdynz/componenty/src/types/ColumnMeta'
import { computed, ref, Ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const columns: Record<string, ColumnMeta> = {
  actions: {
    heading: '',
    format: 'custom',
    sortEnabled: false,
  },

  name: {
    heading: 'Name',
    format: 'text',
    align: 'left',
  },
  username: {
    heading: 'Username',
    format: 'text',
    align: 'left',
  },
  email: {
    heading: 'Email',
    format: 'text',
    align: 'left',
  },
  phone: {
    heading: 'Phone',
    format: 'text',
    align: 'left',
  },
  // role: {
  //   heading: 'Role',
  //   format: 'text',
  //   align: 'left',
  // },
  // initials: {
  //   heading: 'Initials',
  //   format: 'text',
  //   align: 'left',
  // },
  dateCreated: {
    heading: 'Date Created',
    format: 'date',
    align: 'left',
  },
  dateModified: {
    heading: 'Date Modified',
    format: 'date',
    align: 'left',
  },
}

const router = useRouter()
const route = useRoute()
const rows: Ref<Person[]> = ref([])
const pagedInfo: Ref<PagedInfo> = ref({
  orderBy: 'date_created',
  direction: PagedInfo.Direction.ASC,
  pageNumber: 1,
  limit: 25,
  total: 100,
  search: '',
})

const actions = computed(() => {
  return [
    {
      name: 'Create Person',
      iconLeft: 'far-heart',
      action: () => {
        create()
      },
    },
  ]
})

function changePageInfo(pi: PagedInfo) {
  pagedInfo.value = pi
  loadData()
}

async function loadData() {
  const resp = await api.person.pagedPeople(pagedInfo.value)
  pagedInfo.value = resp.pagedInfo
  rows.value = resp.records
}

function create() {
  router.push({ name: 'PersonCreate' })
}

function edit(personUlid: string) {
  router.push({ name: 'PersonEdit', params: { ulid: personUlid } })
}

watch(
  () => route.name,
  (routeName) => {
    if (routeName === 'PersonList') {
      loadData()
    }
  },
  { flush: 'pre', immediate: true, deep: true }
)
</script>
