<template>
  <a-space direction="vertical">
    <div class="buttons">
      <a-button type="primary" @click="handleShowCreate">创建</a-button>
    </div>

    <a-table :dataSource="tagList" :columns="columns">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <span>
            <a @click="handleShowEdit(record.id, record.tag_name)">修改</a>
            <a-divider type="vertical" />
            <a @click="handleDelete(record.id)">删除</a>
          </span>
        </template>
        <template v-else>
          {{ record[column.dataIndex] }}
        </template>
      </template>
    </a-table>

    <a-modal v-model:open="modalVisible" :closable="false" @ok="handleSubmit">
      <a-form :model="formData">
        <a-form-item label="标签名称">
          <a-input v-model:value="formData.name" />
        </a-form-item>
      </a-form>
    </a-modal>
  </a-space>
</template>
<script setup>
import { reactive, ref } from 'vue';
import {
  GetManageTagListApi,
  EditTagApi,
  CreateTagApi,
  DelTagApi,
} from '@/api/tag.js';
import { Modal } from 'ant-design-vue';
import { CreateSuccessMessage, CreateErrorMessage } from '@/utils/alert.js';

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
  },
  {
    title: '名称',
    dataIndex: 'tag_name',
    key: 'tag_name',
  },
  {
    title: '操作',
    key: 'action',
  },
];

const formData = reactive({
  name: '',
});
const modalVisible = ref(false);
const isEdit = ref(false); //false创建，true编辑

// 获取标签
let tagList = ref([]);
const handleGetTagList = async () => {
  const { data } = await GetManageTagListApi();
  tagList.value = data;
};
handleGetTagList();

const handleShowCreate = () => {
  modalVisible.value = true;
  isEdit.value = false;
  formData.id = '';
  formData.name = '';
};

const handleShowEdit = (id, name) => {
  modalVisible.value = true;
  isEdit.value = true;
  formData.id = id;
  formData.name = name;
};

const handleSubmit = () => {
  if (isEdit.value) {
    let params = {
      id: formData.id,
      tag_name: formData.name,
    };

    EditTagApi(params)
      .then((result) => {
        CreateSuccessMessage('编辑成功');
        modalVisible.value = false;
      })
      .catch((err) => {
        console.error(err);
      });
  } else {
    let params = {
      tag_name: formData.name,
    };
    CreateTagApi(params)
      .then((result) => {
        CreateSuccessMessage('创建成功');
        modalVisible.value = false;
      })
      .catch((err) => {
        console.error(err);
      });
  }
};

const handleDelete = (id) => {
  Modal.confirm({
    centered: true,
    content: `您确定要删除该条数据吗？`,
    okText: '确认',
    onOk() {
      DelTagApi({ id: id })
        .then((result) => {
          CreateSuccessMessage('删除成功');

          //   前端模拟实现删除效果
          tagList.value = tagList.value.filter((item) => item.id !== id);
        })
        .catch((err) => {
          console.error(err);
        });
    },
    cancelText: '取消',
    onCancel() {},
  });
};

const showModal = () => {
  open.value = true;
};
const handleOk = (e) => {
  console.log(e);
  open.value = false;
};
</script>

<style lang="scss" scoped></style>
