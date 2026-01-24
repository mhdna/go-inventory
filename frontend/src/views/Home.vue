<script setup>
import { onMounted, ref, watch } from "vue";
import { GetItems } from "../../wailsjs/go/main/App";
import InventoryTable from "../components/InventoryTable.vue";
import CreateItem from "../components/CreateItem.vue";

const loading = ref(true);
const search = ref("");
const items = ref([]);

watch(search, loadItems());

async function loadItems() {
  loading.value = true
  items.value = await GetItems(search.value)
  loading.value = false
}

onMounted(() => loadItems());
</script>

<template>
    <v-container>
        <v-row class="mb-8 mt-1 d-flex align-center">
            <CreateItem @itemCreated="loadItems" />

        </v-row>
        <v-text-field
            label="Search Input"
            v-model="search"
            variant="solo-filled"
            @input="loadItems"
        ></v-text-field>
        <InventoryTable :items="items" :loading="loading" />
    </v-container>
</template>
