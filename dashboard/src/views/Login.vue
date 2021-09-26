<template>
  <div class="home">
    <img alt="UniqueStudio logo" class="logo" src="../assets/logo.jpg" />
    <h1>UniqueStuio SSO System</h1>
    <el-row justify="center">
      <el-col
        :xs="{ span: 20 }"
        :sm="{ span: 14 }"
        :md="{ span: 12 }"
        :lg="{ span: 10 }"
        :xl="{ span: 8 }"
      >
        <Form />
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import Form from "@/components/Form.vue";
import { ElContainer, ElRow, ElCol } from "element-plus";

@Options({
  components: {
    Form,
    ElContainer,
    ElRow,
    ElCol,
  },
})
export default class Login extends Vue {}
</script>

<style scoped>
.logo {
  max-width: 25%;
}
</style>
