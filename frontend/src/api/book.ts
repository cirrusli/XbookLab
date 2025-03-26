import service from './index';

export const fetchBooks = (params: any) => service.get('/books', { params });
export const createBook = (data: any) => service.post('/books', data);
export const updateBook = (id: number, data: any) => service.put(`/books/${id}`, data);
export const deleteBook = (id: number) => service.delete(`/books/${id}`);
export const searchBooks = (keyword: string) => service.get('/books/search', { params: { q: keyword } });