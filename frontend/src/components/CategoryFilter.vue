<template>
  <div class="category-filter">
    <div class="search-box">
      <input 
        type="text"
        placeholder="搜索分类..."
        v-model="searchQuery"
        class="search-input"
      />
    </div>
    
    <div class="category-list">
      <div
        v-for="category in filteredCategories"
        :key="category.id"
        class="category-item"
        :class="{ 'active': selectedCategory === category.id }"
        @click="selectCategory(category.id)"
      >
        {{ category.name }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface Category {
  id: number
  name: string
}

const categories = ref<Category[]>([
  { id: 1, name: '文学小说' },
  { id: 2, name: '科技前沿' },
  { id: 3, name: '历史人文' },
  { id: 4, name: '经济管理' },
  { id: 5, name: '生活艺术' },
  { id: 6, name: '心理成长' },
])

const searchQuery = ref('')
const selectedCategory = ref<number | null>(null)

const filteredCategories = computed(() => {
  return categories.value.filter(c => 
    c.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

const selectCategory = (id: number) => {
  selectedCategory.value = selectedCategory.value === id ? null : id
}
</script>

<style scoped>
.category-filter {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.search-box {
  margin-bottom: 20px;
}

.search-input {
  width: 100%;
  padding: 10px 15px;
  border: 1px solid #e0e0e0;
  border-radius: 24px;
  font-size: 14px;
  outline: none;
}

.search-input:focus {
  border-color: #007bff;
}

.category-list {
  max-height: 500px;
  overflow-y: auto;
}

.category-item {
  padding: 12px 15px;
  margin: 8px 0;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s;
  background: #f8f9fa;
}

.category-item:hover {
  background: #e9ecef;
}

.category-item.active {
  background: #007bff;
  color: white;
  font-weight: 500;
}
</style>