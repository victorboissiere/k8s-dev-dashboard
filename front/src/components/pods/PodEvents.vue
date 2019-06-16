<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <ErrorMessage :message="errorMessage"/>
    <v-dialog v-model="dialog" width="80%">
      <v-card>
        <v-card-title>
          <span class="headline">Events</span>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon @click="refresh">refresh</v-icon>
          </v-btn>
        </v-card-title>
        <v-card-text>
          <Request :path="`namespaces/${namespace}/deployments/${deployment}/events`" ref="request">
            <template v-slot:default="events">
              <div v-if="events.data">
                <div v-if="events.data.length === 0">
                  <p class="text-md-center">No events</p>
                </div>
                  <v-list v-else>
                    <v-list-tile v-for="event in events.data" :key="`${event.name}-${deployment}`">
                      <v-list-tile-content>
                        <v-list-tile-title>
                          {{ event.reason }} ({{ event.count }})
                        </v-list-tile-title>
                        <v-list-tile-sub-title>
                          {{ event.message }}
                        </v-list-tile-sub-title>
                      </v-list-tile-content>
                      <v-list-tile-action>
                        {{ event.lastTimestamp }}
                      </v-list-tile-action>
                    </v-list-tile>
                  </v-list>
              </div>
            </template>
          </Request>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" flat="flat" @click="hide">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import ErrorMessage from '../utils/ErrorMessage.vue';
import Request from '../utils/Request.vue';

export default {
  components: { Request, ErrorMessage },
  props: {
    namespace: { type: String },
    deployment: { type: String },
  },
  methods: {
    refresh() {
      this.$refs.request.refresh();
    },
    show() {
      this.refresh();
      this.dialog = true;
    },
    hide() {
      this.dialog = false;
    },
  },
  data() {
    return {
      isDialogLoading: false,
      dialog: false,
      errorMessage: '',
    };
  },
};
</script>
