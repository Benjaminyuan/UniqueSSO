<template>
  <v-container justify-center fill-height>
    <v-card color="basil">
      <v-card-title class="text-center justify-center py-6">
        <h1 class="font-weight-bold text-h5 basil--text">{{ title }}</h1>
      </v-card-title>

      <v-tabs
        v-model="tab"
        background-color="transparent"
        center-active
        centered
        grow
        color="basil"
      >
        <v-tab v-for="item in items" :key="item.idType">
          {{ item.name }}
        </v-tab>
      </v-tabs>

      <v-tabs-items v-model="tab">
        <v-tab-item v-for="item in items" :key="item.idType">
          <v-card color="basil" flat>
            <Form v-bind:idType="item.idType"></Form>
          </v-card>
        </v-tab-item>
      </v-tabs-items>
    </v-card>
  </v-container>
</template>

<script>
import Form from "../components/Form";
import { IDType } from "../common/const";

export default {
  name: "Login",

  components: {
    Form,
  },

  data() {
    return {
      title: "登录",
      tab: null,
      items: [
        {
          name: "邮箱",
          idType: IDType.Email,
        },
        {
          name: "手机号",
          idType: IDType.Phone,
        },
        {
          name: "短信验证码",
          idType: IDType.SMS,
        },
        {
          name: "企业微信扫码",
          idType: IDType.Wechat,
        },
      ],
    };
  },
};
</script>

<style>
/* Helper classes */
.basil {
  background-color: #fffbe6 !important;
}
.basil--text {
  color: #356859 !important;
}
.v-slide-group__prev {
  display: none !important;
}

.login-form {
  height: 100%;
}
</style>