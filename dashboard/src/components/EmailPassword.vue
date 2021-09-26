<template>
  <el-form ref="form" :model="form" label-width="80px">
    <el-form-item label="邮箱">
      <el-input v-model="_form.email"></el-input>
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="_form.password"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">登陆</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { ElForm, ElFormItem, ElInput, ElButton } from "element-plus";
import { EmailForm } from "@/model/form";
import { RestClient } from "@/api/rest";

@Options({
  props: {},
  components: {
    ElForm,
    ElFormItem,
    ElInput,
    ElButton,
  },
})
export default class EmailPassword extends Vue {
  // HACK: I don't know the reason why the virable can't be form.
  // Once I used form as it, I can't enter char in the input box.
  private _form = {
    email: "",
    password: "",
  };

  private client = new RestClient("http://localhost:5000");

  public set form(f: EmailForm) {
    this._form.email = f.Email;
    this._form.password = f.Password;
  }

  public get form(): EmailForm {
    return {
      Email: this._form.email,
      Password: this._form.password,
    };
  }

  async onSubmit(): Promise<void> {
    console.log(this._form);
    console.log(this.form);
    try {
      const resp = await this.client.loginByEmail(
        this.form,
        "https://bing.com"
      );
      console.log(resp);
    } catch (err) {
      console.log(err);
    }
  }
}
</script>
