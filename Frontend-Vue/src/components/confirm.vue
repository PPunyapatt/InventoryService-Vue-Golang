<template>
    <div class="text-center pa-4">
    <v-dialog
      v-model="dialog"
      max-width="400"
      persistent
    >
      <template v-slot:activator="{ props: activatorProps }">
        <v-btn v-bind="activatorProps">
          Open Dialog
        </v-btn>
      </template>

      <v-card
        prepend-icon="mdi-map-marker"
        text="Let Google help apps determine location. This means sending anonymous location data to Google, even when no apps are running."
      >
        <template v-slot:actions>
          <v-spacer></v-spacer>

          <v-btn @click="dialog = false">
            Disagree
          </v-btn>

          <v-btn @click="dialog = false">
            Agree
          </v-btn>
        </template>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
export default {
  props: {
    open: { type: Boolean, required: true },
    text: { type: String, required: true }
  },
  setup(props) {
    const dialog = ref(false)
    return { dialog }
  }
}
</script>
