<script setup>
import { onMounted, ref, watch } from 'vue'
import { GetItems } from '../../wailsjs/go/main/App'
import InventoryTable from '../components/InventoryTable.vue'
import CreateItem from '../components/CreateItem.vue'

const loading = ref(true)
const search = ref('')
const items = ref([])

async function loadItems() {
  loading.value = true
  items.value = await GetItems(search.value)
  loading.value = false
}

onMounted(() => loadItems())
</script>

<template>
  <v-container>
    <InventoryTable :items="items" :loading="loading" />
  </v-container>
</template>

