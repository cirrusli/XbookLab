<template>
  <div class="form-group-wrap">
    <!-- <div class="form-group">
      <label for="" class="label">账号：</label>
      <input type="text" v-model="formData.email" class="form-input" readonly />
    </div> -->
    <div class="form-group">
      <label for="" class="label">用户名：</label>
      <input
        type="text"
        v-model="formData.name"
        class="form-input"
        placeholder="请输入昵称"
      />
    </div>
    <div class="form-group">
      <label for="" class="label">头像：</label>
      <img :src="formData.avatar" alt="" class="avatar" />
    </div>

    <div class="form-footer">
      <button class="submit-btn sm" @click="handleUpdateUserInfo">保存</button>
      <button class="btn-default sm">
        <router-link to="/user">返回个人主页</router-link>
      </button>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { UpdateUserInfoApi } from '@/api/user.js';
import { CreateSuccessMessage, CreateErrorMessage } from '@/utils/alert.js';
import { useRouter } from 'vue-router';

const router = useRouter();

const formData = reactive({
  email: '',
  name: '',
  avatar: '',
});

const getUserInfo = (() => {
  // 模拟获取用户信息
  const userInfo = JSON.parse(localStorage.getItem('userInfo'));
  formData.email = userInfo.email;
  formData.name = userInfo.name;
  formData.avatar = userInfo.avatar;
})();

const handleUpdateUserInfo = async () => {
  if (formData.name === '') {
    CreateErrorMessage('昵称不能为空');
    return false;
  }
  // 模拟保存用户信息
  UpdateUserInfoApi(formData).then((res) => {
    if (res.code === 200) {
      localStorage.setItem('userInfo', JSON.stringify(formData));
      CreateSuccessMessage('保存成功');
      router.push('/user');
    } else {
      CreateErrorMessage('保存失败');
    }
  });
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
