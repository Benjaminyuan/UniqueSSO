<template>
  <el-card class="box-card">
    <h2>登陆</h2>
    <el-tabs v-model="activeName" @tab-click="handleClick" stretch="true">
      <el-tab-pane label="邮箱" name="email">
        <EmailPassword />
      </el-tab-pane>
      <el-tab-pane label="手机号" name="phone">
        <PhonePassword />
      </el-tab-pane>
      <el-tab-pane label="验证码" name="sms">
        <PhoneSMS />
      </el-tab-pane>
      <el-tab-pane label="Lark扫码" name="lark">
        <Oauth />
      </el-tab-pane>
    </el-tabs>
  </el-card>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { ElCard, ElButton, ElTabs, ElTabPane } from "element-plus";
import EmailPassword from "@/components/EmailPassword.vue";
import PhonePassword from "@/components/PhonePassword.vue";
import PhoneSMS from "@/components/PhoneSMS.vue";
import Oauth from "@/components/Oauth.vue";

@Options({
  props: {
    msg: String,
  },
  components: {
    ElCard,
    ElButton,
    ElTabs,
    ElTabPane,
    EmailPassword,
    PhonePassword,
    PhoneSMS,
    Oauth,
  },
})
export default class Form extends Vue {
  msg!: string;
}
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  /* max-width: 100%; */
  background-color: #fffbe6 !important;
}
</style>
