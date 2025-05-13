<template>
  <a-space direction="vertical">
    <div class="buttons">
      <a-button type="primary" @click="goToCreate">创建</a-button>
    </div>

    <a-table :dataSource="bookList" :columns="columns">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'cover'">
          <img
            :src="record.cover"
            alt="封面"
            style="width: 50px; height: 50px"
          />
        </template>
        <template v-else-if="column.key === 'action'">
          <span>
            <!-- <a @click="handleEdit(record.id)">修改</a>
            <a-divider type="vertical" /> -->
            <a @click="handleDelete(record.id)">删除</a>
          </span>
        </template>
        <template v-else>
          {{ record[column.dataIndex] }}
        </template>
      </template>
    </a-table>
  </a-space>
</template>
<script setup>
import { ref } from 'vue';
import { GetBooksListApi } from '@/api/recommend.js';
import { useRouter } from 'vue-router';
import { Modal } from 'ant-design-vue';

const router = useRouter();

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
  },
  {
    title: '标题',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: '作者',
    dataIndex: 'author',
    key: 'author',
  },
  {
    title: '评分',
    dataIndex: 'rating',
    key: 'rating',
  },
  {
    title: '封面',
    dataIndex: 'cover',
    key: 'cover',
  },
  {
    title: '操作',
    key: 'action',
  },
];

// 获取书单
let bookList = ref([]);
const handleGetBookList = async () => {
  const { data } = await GetBooksListApi();
  bookList.value = data.Books.map(book => ({
    id: book.BookID,
    title: book.Title,
    author: book.Author,
    cover: book.Cover,
    rating: book.AverageRating,
    desc: book.Description,
    tag: book.Category
  }));
};
handleGetBookList();

const goToCreate = () => {
  // 跳转到创建书籍的页面
  router.push('/manage/book/create');
};

const handleEdit = (id) => {
  router.push({ path: `/manage/book/edit/${id}` });
};

const handleDelete = (id) => {
  Modal.confirm({
    centered: true,
    content: `您确定要删除该条数据吗？`,
    okText: '确认',
    onOk() {
      bookList.value = bookList.value.filter((item) => item.id !== id);
    },
    cancelText: '取消',
    onCancel() {},
  });
};
</script>

<style lang="scss" scoped></style>
