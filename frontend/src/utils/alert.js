import { message } from 'ant-design-vue'

// 成功消息提示
export function CreateSuccessMessage(msg) {
    message.config({
        top: `100px`,
        duration: 3,
        maxCount: 3
    })
    message.success(msg)
}

// 失败消息提示
export function CreateErrorMessage(msg) {
    message.config({
        top: `100px`,
        duration: 3,
        maxCount: 3
    })
    message.error({
        content: msg,
        style: {
            color: '#ff0000',
        }
    })
}