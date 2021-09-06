<template>
  <v-row justify="center">
    <v-col>
      <div v-if="idType === 'wechat'">
        <v-card>
          <v-container justify-center fill-height>
            <v-btn
              class="ma-2"
              v-show="!showQrcode"
              color="secondary"
              @click="getQrcode"
              >获取企业微信二维码
            </v-btn>
            <v-img
              v-show="showQrcode"
              max-height="260"
              max-width="260"
              :src="qrcodeSrc"
            >
            </v-img>
          </v-container>
        </v-card>
      </div>

      <div v-else>
        <v-card ref="form">
          <div v-if="idType === 'email'">
            <v-card-text>
              <v-text-field
                ref="id"
                v-model="id"
                :error-messages="errorMessages"
                label="邮箱"
                :rules="emailRules"
                required
              ></v-text-field>
              <v-text-field
                ref="password"
                v-model="password"
                label="密码"
                type="password"
                required
              ></v-text-field>
            </v-card-text>
          </div>

          <div v-else-if="idType === 'phone'">
            <v-card-text>
              <v-text-field
                ref="id"
                v-model="id"
                :error-messages="errorMessages"
                label="手机号"
                :rules="phoneRules"
                prefix="+86"
                required
              ></v-text-field>
              <v-text-field
                ref="password"
                v-model="password"
                label="密码"
                type="password"
                required
              ></v-text-field>
            </v-card-text>
          </div>

          <div v-else-if="idType === 'sms'">
            <v-card-text>
              <v-text-field
                ref="id"
                v-model="id"
                :error-messages="errorMessages"
                label="手机号"
                :rules="phoneRules"
                prefix="+86"
                required
              ></v-text-field>
              <v-row>
                <v-col cols="7">
                  <v-text-field
                    ref="password"
                    v-model="password"
                    label="验证码"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="5">
                  <v-btn
                    class="ma-2"
                    :loading="loading"
                    :disabled="smsRemainSeconds > 0"
                    color="secondary"
                    @click="smsCountDown()"
                  >
                    {{ smsButtonText }}
                  </v-btn>
                </v-col>
              </v-row>
            </v-card-text>
          </div>

          <v-divider class="mt-12"></v-divider>
          <v-card-actions>
            <v-btn text> Cancel </v-btn>
            <v-spacer></v-spacer>
            <v-slide-x-reverse-transition>
              <v-tooltip v-if="formHasErrors" left>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    icon
                    class="my-0"
                    v-bind="attrs"
                    @click="resetForm"
                    v-on="on"
                  >
                    <v-icon>mdi-refresh</v-icon>
                  </v-btn>
                </template>
                <span>Refresh form</span>
              </v-tooltip>
            </v-slide-x-reverse-transition>
            <v-btn color="primary" text @click="submit"> Submit</v-btn>
          </v-card-actions>
        </v-card>
      </div>
    </v-col>
  </v-row>
</template>

<script>
import { Login, FetchSmsCode, FetchQrCodeSrc } from "../common/api";

export default {
  name: "Form",
  data: () => ({
    id: null,
    password: null,
    formHasErrors: false,
    errorMessages: "",
    loginForm: {},

    areaCode: ["+86"],

    qrcodeSrc: null,
    showQrcode: false,

    loading: false,
    smsButtonText: "发送",
    smsDuration: 60,
    smsRemainSeconds: 0,

    emailRules: [
      (v) => !!v || "邮箱不能为空",
      (v) => /.+@.+\..+/.test(v) || "邮箱格式错误",
    ],
    phoneRules: [
      (v) => !!v || "手机号不能为空",
      (v) => /^1[0-9]{10}$/.test(v) || "邮箱格式错误",
    ],
  }),

  computed: {
    form() {
      return {};
    },
  },

  props: {
    idType: String,
  },

  watch: {
    id() {
      this.errorMessages = "";
    },
  },

  methods: {
    resetForm() {
      this.errorMessages = "";
      this.formHasErrors = false;

      this.$refs["id"]?.reset();
      this.$refs["password"]?.reset();
    },

    async smsCountDown() {
      this.smsRemainSeconds = this.smsDuration;
      this.loading = true;
      // TODO: set fetch API

      try {
        const msg = await FetchSmsCode(this.id);
        this.$toasted.show(msg, { type: "success" });
      } catch (err) {
        this.$toasted.show(err, { type: "error" });
      }

      this.loading = false;

      // count down
      let countDown = setInterval(() => {
        this.smsRemainSeconds--;
        if (this.smsRemainSeconds <= 0) {
          this.smsButtonText = "发送";
          clearInterval(countDown);
        } else {
          this.smsButtonText = `发送(${this.smsRemainSeconds}秒)`;
        }
      }, 1000);
    },

    async getQrcode() {
      try {
        const imgsrc = await FetchQrCodeSrc();
        this.qrcodeSrc = imgsrc;
        this.showQrcode = true;

        const redirectUrl = await Login(
          this.idType,
          this.qrcodeSrc,
          this.password,
          this.$route.query.service || ""
        );
        window.location = redirectUrl;
      } catch (e) {
        this.$toasted.show(e, { type: "error" });
      } finally {
        this.showQrcode = false;
      }
    },

    async submit() {
      this.formHasErrors = false;

      Object.keys(this.form).forEach((f) => {
        if (!this.form[f]) this.formHasErrors = true;
        if (!this.$refs[f].validate()) this.formHasErrors = true;
      });

      if (this.formHasErrors) {
        this.$toasted.show("表单格式错误");
        return;
      }

      try {
        const url = await Login(
          this.idType,
          this.id,
          this.password,
          this.$route.query.service || ""
        );
        window.location = url;
      } catch (e) {
        this.$toasted.show(e, { type: "error" });
      }

      this.resetForm();
    },
  },
};
</script>