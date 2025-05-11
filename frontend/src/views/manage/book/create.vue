<template>
  <a-form
    ref="formRef"
    :model="formData"
    :label-col="labelCol"
    :wrapper-col="wrapperCol"
    :rules="rules"
    class="section-box"
  >
    <a-form-item label="标题" name="title">
      <a-input v-model:value="formData.title" />
    </a-form-item>
    <a-form-item label="作者" name="author">
      <a-input v-model:value="formData.author" />
    </a-form-item>
    <a-form-item label="评分" name="rating">
      <a-input-number
        id="inputNumber"
        v-model:value="formData.rating"
        :min="1"
        :max="10"
        :precision="1"
        :step="0.1"
      />
    </a-form-item>
    <a-form-item label="封面" name="cover">
      <a-upload
        v-model:file-list="state.fileList"
        list-type="picture-card"
        class="avatar-uploader"
        :show-upload-list="false"
        :accept="['image/jpeg', 'image/png', 'image/jpg']"
        :before-upload="handleBeforeUpload"
        @change="handleUploadChange"
        :custom-request="handleUploadImg"
      >
        <img v-if="formData.cover" :src="formData.cover" alt="封面" />
        <div v-else>
          <div class="ant-upload-text">上传</div>
        </div>
      </a-upload>
    </a-form-item>
    <a-form-item label="简介" name="desc">
      <a-textarea v-model:value="formData.desc" />
    </a-form-item>
    <a-form-item label="标签" name="category">
      <a-radio-group v-model:value="formData.category">
        <a-radio
          v-for="(item, index) in state.tagList"
          :key="index"
          :value="item.tag_name"
          >{{ item.tag_name }}</a-radio
        >
      </a-radio-group>
    </a-form-item>
    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
      <a-space :size="10">
        <a-button type="primary" @click="onSubmit">保存</a-button>
        <a-button @click="handleCancel">取消</a-button>
      </a-space>
    </a-form-item>
  </a-form>
</template>
<script setup>
import { ref, reactive } from 'vue';
import { CreateSuccessMessage, CreateErrorMessage } from '@/utils/alert.js';
import { CreateBookApi, UploadImgApi } from '@/api/recommend.js';
import { useRouter } from 'vue-router';
import { GetManageTagListApi } from '@/api/tag.js';

const router = useRouter();

const formRef = ref({});
const formData = reactive({
  title: '',
  author: '',
  rating: '',
  desc: '',
  category: '',
  cover: '',
});

const state = reactive({
  fileList: [],
  tagList: [],
});

const handleGetTagList = async () => {
  const { data } = await GetManageTagListApi();
  state.tagList = data;
};
handleGetTagList();

const rules = {
  title: [
    { required: true, message: '请输入书籍标题', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' },
  ],
  author: [
    { required: true, message: '请输入书籍作者', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' },
  ],
  rating: [
    { required: true, message: '请输入书籍评分', trigger: 'blur' },
    {
      type: 'number',
      min: 1,
      max: 10,
      message: '评分范围在1到10之间',
      trigger: 'blur',
    },
  ],
  desc: [
    { required: true, message: '请输入书籍简介', trigger: 'blur' },
    { min: 10, max: 200, message: '长度在 10 到 200 个字符', trigger: 'blur' },
  ],
  category: [
    { required: true, message: '请选择书籍标签', trigger: 'change' },
    { type: 'string', message: '请选择书籍标签', trigger: 'change' },
  ],
  cover: [
    { required: true, message: '请上传书籍封面', trigger: 'change' },
    { type: 'string', message: '请上传书籍封面', trigger: 'change' },
  ],
};

const onSubmit = () => {
  console.log(formData);
  formRef.value
    .validate()
    .then(() => {
      CreateBookApi(formData)
        .then((res) => {
          if (res.code === 200) {
            CreateSuccessMessage('创建成功');
            router.push('/manage/book');
          } else {
            CreateErrorMessage(res.msg);
          }
        })
        .catch((error) => {
          console.log('error', error);
        });
    })
    .catch((error) => {
      console.log('error', error);
    });
};
const labelCol = {
  style: {
    width: '150px',
  },
};
const wrapperCol = {
  span: 14,
};

const handleBeforeUpload = (file) => {
  return new Promise((resolve) => {
    const isJpgOrPng =
      file.type === 'image/jpeg' ||
      file.type === 'image/png' ||
      file.type === 'image/jpg';
    if (!isJpgOrPng) {
      CreateErrorMessage('图片格式只能为jpeg/png/jpg');
      return false;
    }
    const isLt3M = file.size / 1024 / 1024 <= 3;
    if (!isLt3M) {
      CreateErrorMessage('图片大小不得超过3M');
      return false;
    }
    return resolve(true);
  });
};

// 自定义上传图片
const handleUploadImg = async (data) => {
  const { file, onProgress, onSuccess } = data;
  let newFormData = new FormData();
  newFormData.append('feed_img', file);
  UploadImgApi(newFormData)
    .then((res) => {
      if (res.code === 200) {
        formData.cover = res.data.url;
        CreateSuccessMessage('上传成功');
      } else {
        CreateErrorMessage(res.msg);
      }
    })
    .catch((error) => {
      console.log('error', error);
    });
};

const handleUploadChange = ({ file, fileList }) => {
  // state.fileList = fileList
};

const handleCancel = () => {
  formData.title = '';
  formData.author = '';
  formData.rating = '';
  formData.desc = '';
  formData.category = '';
  formData.cover = '';
  router.push('/manage/book');
};
</script>
