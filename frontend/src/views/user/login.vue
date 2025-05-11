<template>
  <div class="height-full flex aic jcc">
    <div class="modal">
      <div class="modal-header">
        <h3 class="modal-title">欢迎来到XbookLab!</h3>
      </div>
      <div class="modal-body">
        <!-- <div class="social-login">
          <button class="social-btn google">
            <img
              src="https://cdn.jsdelivr.net/gh/googlefonts/noto-emoji/svg/emoji_u1f310.svg"
              width="20"
              alt="Google"
            />
            使用Google登录
          </button>
          <button class="social-btn wechat">
            <img
              src="https://cdn.jsdelivr.net/gh/googlefonts/noto-emoji/svg/emoji_u1f4e1.svg"
              width="20"
              alt="WeChat"
            />
            使用微信登录
          </button>
        </div> -->
        <!-- <div class="divider">
          <span>或</span>
        </div> -->
        <form id="login-form">
          <div class="form-group">
            <label for="email" class="form-label">电子邮箱</label>
            <input
              type="email"
              v-model="formData.email"
              class="form-input"
              placeholder="请输入您的邮箱"
            />
          </div>
          <div class="form-group">
            <label for="password" class="form-label">密码</label>
            <input
              type="password"
              v-model="formData.password"
              class="form-input"
              placeholder="请输入密码"
            />
          </div>
          <div class="form-footer">
            <router-link to="/forgotPwd" class="forgot-password"
              >忘记密码?</router-link
            >
          </div>
        </form>
      </div>
      <div class="modal-footer">
        <button class="submit-btn" @click="handleLogin">登录</button>
        <p class="signup-text">
          还没有账号? <router-link to="/register">立即注册</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { checkEmail } from '@/utils/utils';
import { CreateErrorMessage, CreateSuccessMessage } from '@/utils/alert.js';
import { LoginApi } from '@/api/user.js';
import { useRouter } from 'vue-router';

const router = useRouter();

const formData = reactive({
  email: 'admin@gmail.com',
  password: '123456',
});

const validateForm = () => {
  const { email, password } = formData;
  if (!email) {
    CreateErrorMessage('请输入邮箱');
    return false;
  }
  if (!checkEmail(email)) {
    CreateErrorMessage('请输入有效的邮箱地址');
    return false;
  }
  if (!password) {
    CreateErrorMessage('请输入密码');
    return false;
  }
  if (password.length < 6) {
    CreateErrorMessage('密码长度至少为6位');
    return false;
  }
  return true;
};

const handleLogin = () => {
  const validate_res = validateForm();
  if (!validate_res) {
    return;
  }

  LoginApi(formData)
    .then((result) => {
      CreateSuccessMessage('登录成功');
      const { data } = result;
      localStorage.setItem('token', data.token);
      localStorage.setItem('userInfo', JSON.stringify(data));
      router.push('/');
    })
    .catch((err) => {
      console.error(err);
    });
};
</script>

<style lang="scss" scoped>
.modal {
  background: white;
  border-radius: 12px;
  width: 400px;
  max-width: 90%;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  transition: var(--transition);
}

.modal-header {
  padding: 20px;
  position: relative;
}

.modal-title {
  font-size: 22px;
  font-weight: 600;
  text-align: center;
  color: var(--primary-color);
}

.modal-body {
  padding: 20px;
}

.form-group {
  margin-bottom: 16px;
}

.form-label {
  display: block;
  margin-bottom: 8px;
  font-size: 14px;
  color: var(--text-dark);
}

.form-input {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  font-size: 14px;
  transition: var(--transition);
  box-sizing: border-box;
  &:focus {
    border-color: var(--primary-color);
    outline: none;
    box-shadow: 0 0 0 2px rgba(66, 185, 131, 0.2);
  }
}

.modal-footer {
  padding: 0 20px 20px;
  display: flex;
  flex-direction: column;
  /* 改为垂直布局 */
  align-items: center;
  /* 居中显示 */
  gap: 15px;
  /* 添加间距 */
}

.submit-btn {
  background-color: var(--primary-color);
  color: white;
  border: none;
  padding: 10px 24px;
  border-radius: 6px;
  font-weight: 500;
  cursor: pointer;
  transition: var(--transition);
  width: auto;
  /* 改为自动宽度 */
  min-width: 200px;
  /* 设置最小宽度 */
  &:hover {
    background-color: #3aa876;
  }
}

.form-footer {
  text-align: right;
}

.forgot-password {
  font-size: 13px;
  color: var(--primary-color);
}

.signup-text {
  margin-top: 15px;
  font-size: 14px;
  color: var(--text-light);
  text-align: center;
  a {
    color: var(--primary-color);
    font-weight: 500;
  }
}
</style>
