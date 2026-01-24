<template>
    <v-btn @click="dialog = true">Create Item</v-btn>

    <v-dialog v-model="dialog" width="700" persistent>
        <v-sheet class="mx-auto" width="700" min-height="550">
            <v-form validate-on="submit lazy" @submit.prevent="submit">
                <v-text-field
                    v-model="code"
                    label="Product Code"
                ></v-text-field>

                <v-text-field
                    v-model="name"
                    label="Product Name"
                ></v-text-field>

                <v-text-field
                    v-model="description"
                    label="Product Description"
                ></v-text-field>
                <v-text-field
                    v-model.number="quantity"
                    label="Initial Quantity"
                    type="number"
                ></v-text-field>

                <v-btn
                    :disabled="loading"
                    class="mt-2"
                    text="Cancel"
                    block
                    @click="cancel"
                ></v-btn>

                <v-btn
                    :loading="loading"
                    class="mt-2"
                    text="Submit"
                    type="submit"
                    block
                ></v-btn>
                {{ error }}
            </v-form>
        </v-sheet>
    </v-dialog>
</template>

<script setup>
import { ref } from "vue";
import { CreateItem, Print } from "../../wailsjs/go/main/App";

// to update the table in parent
const emit = defineEmits(["itemCreated"]);

const dialog = ref(false);
const loading = ref(false);
const code = ref("");
const name = ref("");
const description = ref("");
const quantity = ref(0);

const error = ref("");

async function submit(event) {
    const results = await event;
    if (!results.valid) return;

    loading.value = true;
        await CreateItem(
            code.value,
            name.value,
            description.value,
            quantity.value,
        );
        emit("itemCreated");
        dialog.value = false;
        code.value = "";
        name.value = "";
        description.value = "";
        quantity.value = 0;
}

function cancel() {
    dialog.value = false;
    code.value = "";
    name.value = "";
    description.value = "";
    quantity.value = 0;
}
</script>
