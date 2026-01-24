<template>
    <v-toolbar title="Go Inventory System" color="orange">
        <v-toolbar-items>
            <!-- <v-btn>Resources</v-btn> -->
            <v-btn to="/inventory">Inventory</v-btn>
            <v-btn to="/products">Products</v-btn>
            <v-divider class="mx-2" vertical></v-divider>
            <v-btn>About</v-btn>
            <!-- <v-btn icon="mdi-dots-vertical"></v-btn> -->
            <v-btn append-icon="mdi-export"> logout </v-btn>
        </v-toolbar-items>
    </v-toolbar>

    <router-view />
</template>

<style>
h1 {
    text-align: center;
    margin: 20px;
}
</style>
