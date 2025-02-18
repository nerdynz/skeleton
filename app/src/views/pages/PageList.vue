<template>
  <with-actions class="px-4 py-3" :actions="actions" :paged-info="pagedInfo" @pagination-changed="changePageInfo">
    <template #left>
      <h1 class="title is-2 mo-0">Pages</h1>
    </template>
    <template #right>
      <form-field for="search" icon="free-search" label="Search" class="with-button">
        <input class="input" name="search" v-model.trim="pagedInfo.search" type="text" @keyup.enter="loadData" @keyup.esc="pagedInfo.search = ''" autocomplete="off" />
      </form-field>
      <!--<button class="button is-primary" @click="actions.find((a:any) => a.name === 'Create Page')?.action">Create Page</button>-->
    </template>
    <!-- <data-table class="is-fullwidth is-stripped" :columns="columns" :rows="rows" :paged-info="pagedInfo" hide-pagination>
      <template #colgroup>
        <colgroup>
          <col width="100px" />
          <col width="200px" />
          <col width="100px" />
          <col width="auto" />
          <col width="140px" />
          <col width="140px" />
        </colgroup>
      </template>
      <template #actions="{ row }">
        <o-button tag="a" @click="edit(row.pageUlid)" class="is-light" size="small" icon-pack="free" icon-left="edit is-small">Edit</o-button>
      </template>
      <loading :on="['Load']" />
    </data-table> -->
    <o-table
      datakey="pageUlid"
      :data="rows"
      paginated
      :per-page="pagedInfo.limit"
      :backend-pagination="true"
      :backend-sorting="true"
      :total="pagedInfo.total"
      :default-sort="['dateCreated', 'desc']"
      :striped="true"
      :mobile-cards="false"
    >
      <o-table-column v-slot="{row}" field="action">
        <div class="field has-addons">
          <div class="control">
            <RouterLink :to="{name: 'PageEdit', params: {ulid: row.pageUlid} }" class="button is-primary is-small">Edit</RouterLink>
          </div>
        </div>
      </o-table-column>
      <o-table-column v-slot="{row}" field="name" label="Name" sortable>
        {{ row.title }}
      </o-table-column>
      <o-table-column v-slot="{row}" field="status" label="Status" sortable>
        {{ row.summary }}
      </o-table-column>
    </o-table>
  </with-actions>
  <router-view />
</template>

<script lang="ts" setup>
import { PagedInfo } from '@/api/pb/base.pb'
import { Page, PagedPages } from '@/api/pb/page.pb'
import { computed, ref, watch, type Ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const columns: Record<string, any> = {
  actions: {
    heading: '',
    format: 'custom',
    sortEnabled: false,
  },

  // pageUlid: {
  //   heading: 'Page Ulid',
  //   format:  'text',
  //   align: 'left',
  // },
  title: {
    heading: 'Title',
    format:  'text',
    align: 'left',
  },
  content: {
    heading: 'Content',
    format:  'text',
    align: 'left',
  },
  dateCreated: {
    heading: 'Date Created',
    format:  'date',
    align: 'left',
  },
  dateModified: {
    heading: 'Date Modified',
    format:  'date',
    align: 'left',
  },
  // siteUlid: {
  //   heading: 'Site Ulid',
  //   format:  'text',
  //   align: 'left',
  // },
}

const route = useRoute()
const router = useRouter()

let rows: Page[] = $ref([])
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
      name: 'Create Page',
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
  const resp = await PagedPages(pagedInfo.value)
  pagedInfo.value = resp.pagedInfo
  rows = resp.records
}

function create() {
  router.push({ name: 'PageCreate' })
}

function edit(pageUlid: string) {
  router.push({ name: 'PageEdit', params: { ulid: pageUlid } })
}

watch(
  () => route.name,
  (routeName) => {
    if (routeName === 'PageList') {
      loadData()
    }
  },
  { flush: 'pre', immediate: true, deep: true }
)


</script>
