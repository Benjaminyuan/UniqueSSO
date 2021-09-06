import axios from "axios";
import { IDType, BackendUrl } from "./const";

const apiInstance = axios.create({
    baseURL: BackendUrl,
    // timeout: 1000,
    headers: {
        "Content-Type": "application/json"
    }
})

export async function Login(idType, id, pwd, service) {
    let form = {}
    switch (idType) {
        case IDType.Email:
            form = { "email": id, "password": pwd }
            break
        case IDType.Phone:
            form = { "phone": id, "password": pwd }
            break
        case IDType.SMS:
            form = { "phone": `+86${id}`, "code": pwd }
            break
        case IDType.Wechat:
            form = { "qrcode_src": id }
            break
        default:
            throw Error(`unsupported login type: ${idType}`)
    }
    try {
        const resp = await apiInstance.post(`/cas/login?type=${idType}&service=${service}`, form)
        return resp.data?.serviceResponse?.authenticationSuccess?.redirectService
    } catch (err) {
        const resp = err.response
        if (resp.status === 401) throw Error(`${resp.data?.serviceResponse?.authenticationFailure?.description}`);
    }
}

export async function FetchSmsCode(phone) {
    try {
        await apiInstance.post(`/sms/code`, { "phone": `+86${phone}` })
    } catch (err) {
        const resp = err.response
        if (resp.status !== 200) throw Error(`获取验证码失败: ${resp.data?.serviceResponse?.authenticationFailure?.description}`)
    }
    return "验证码已发送"
}

export async function FetchQrCodeSrc() {
    try {
        const resp = await apiInstance.get(`/qrcode/code`)
        return resp.data?.serviceResponse?.authenticationSuccess?.qrcodeSrc
    } catch (err) {
        const resp = err.response
        if (resp.status !== 200) throw Error(`获取企业微信失败: ${resp.data?.serviceResponse?.authenticationFailure?.description}`)
    }
}