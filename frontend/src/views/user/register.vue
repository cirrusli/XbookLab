<template>
  <div class="height-full flex aic jcc">
    <div class="modal">
      <div class="modal-header">
        <h3 class="modal-title">欢迎来到XbookLab!</h3>
      </div>
      <div class="modal-body">
        <form id="login-form">
          <div class="form-group">
            <label for="username" class="form-label">用户名</label>
            <input
              type="username"
              v-model="formData.username"
              class="form-input"
              placeholder="请输入您的用户名"
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
        </form>
      </div>
      <div class="modal-footer">
        <button class="submit-btn" @click="handleRegister">注册</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { checkEmail } from '@/utils/utils';
import { CreateErrorMessage, CreateSuccessMessage } from '@/utils/alert.js';
import { RegisterApi } from '@/api/user.js';
import { useRouter } from 'vue-router';

const router = useRouter();

const formData = reactive({
  username: '',
  password: '123456',
});

const validateForm = () => {
  const { username, password } = formData;
  if (!username) {
    CreateErrorMessage('请输入用户名');
    return false;
  }
  // if (!checkEmail(username)) {
  //   CreateErrorMessage('请输入有效的邮箱地址');
  //   return false;
  // }
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

const handleRegister = () => {
  const validate_res = validateForm();
  if (!validate_res) {
    return;
  }

  RegisterApi(formData)
    .then((result) => {
      CreateSuccessMessage('注册成功，请登录');
      //   localStorage.setItem('token', data.token);
      //   localStorage.setItem('userInfo', JSON.stringify(data));

      //   如果是注册后自动登录的话，就不需要这一步了
      router.push('/login');
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
