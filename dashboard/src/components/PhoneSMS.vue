<template>
  <div></div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";

@Options({
  props: {},
  components: {},
})
export default class PhoneSMS extends Vue {
  msg!: string;
}
</script>
