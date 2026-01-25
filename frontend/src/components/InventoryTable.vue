<template>
  <div>
    <v-data-table
      :items="items"
      :loading="loading"
      fixed-header
      height="520"
      no-data-text="No Items Created Yet"
      >
    </v-data-table>
  </div>
</template>

<script setup>
defineProps({
  items: Array,
  loading: Boolean
})
</script>
