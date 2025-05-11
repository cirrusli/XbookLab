<template>
  <div class="form-group-wrap">
    <div class="form-group">
      <label for="" class="label">旧密码：</label>
      <input
        type="password"
        v-model="formData.old_password"
        class="form-input"
        placeholder="请输入旧密码"
      />
    </div>
    <div class="form-group">
      <label for="" class="label">新密码：</label>
      <input
        type="password"
        v-model="formData.new_password"
        class="form-input"
        placeholder="请输入新密码"
      />
    </div>

    <div class="form-footer">
      <button class="submit-btn sm" @click="handleSubmit">保存</button>
      <button class="btn-default sm">
        <router-link to="/user">返回个人主页</router-link>
      </button>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { UpdateUserInfoApi, ChangePasswordApi } from '@/api/user.js';
import { CreateSuccessMessage, CreateErrorMessage } from '@/utils/alert.js';
import { useRouter } from 'vue-router';

const router = useRouter();

const formData = reactive({
  old_password: '',
  new_password: '',
});

const validateForm = () => {
  const { old_password, new_password } = formData;
  if (!old_password) {
    CreateErrorMessage('请输入旧密码');
    return false;
  }
  if (!new_password) {
    CreateErrorMessage('请输入新密码');
    return false;
  }
  if (old_password.length < 6 || new_password.length < 6) {
    CreateErrorMessage('密码长度至少为6位');
    return false;
  }
  if (old_password === new_password) {
    CreateErrorMessage('新密码不能与旧密码相同');
    return false;
  }
  return true;
};

const handleSubmit = async () => {
  const validate_res = validateForm();
  if (!validate_res) {
    return;
  }

  ChangePasswordApi(formData).then((res) => {
    if (res.code === 200) {
      CreateSuccessMessage('修改成功');
      handleLogout();
      router.push('/login');
    } else {
      CreateErrorMessage('修改失败');
    }
  });
};

const handleLogout = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('userInfo');
};
</script>

<style lang="scss" scoped>
.form-group-wrap {
  background: white;
  border-radius: 8px;
  padding: 20px 40px;
  box-shadow: var(--card-shadow);
}

.form-group {
  display: flex;
  align-items: center;
  margin-bottom: 30px;
}

.label {
  width: 80px;
  font-size: 14px;
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

.form-footer {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 60px;
  button {
    margin-right: 40px;
  }
}

.avatar {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: 50%;
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
  .sm {
    padding: 0 20px;
  }
  /* 设置最小宽度 */
  &:hover {
    background-color: #3aa876;
  }
}
</style>
