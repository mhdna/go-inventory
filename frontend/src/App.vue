<template>
    <v-toolbar title="Go Inventory System" color="orange">
        <v-toolbar-items>
            <v-btn>Dashboard</v-btn>

            <!-- <v-btn>Resources</v-btn> -->
            <v-btn>About</v-btn>
            <v-divider class="mx-2" vertical></v-divider>
            <v-btn>logout</v-btn>
            <!-- <v-btn icon="mdi-dots-vertical"></v-btn> -->
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
