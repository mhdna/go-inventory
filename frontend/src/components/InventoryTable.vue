<template>
  <div>
    <v-data-table
      :items="items"
      :loading="loading"
      fixed-header
      height="520">
    </v-data-table>
  </div>
</template>

<script setup>
defineProps({
  items: Array,
  loading: Boolean
})
</script>
