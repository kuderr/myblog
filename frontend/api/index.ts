import axios from 'axios'

// const API_URL = "http://localhost:5000"; // dev
const API_URL = 'https://kuderblog.com/api' // prod

export function authenticate(userData) {
  return axios.post(`${API_URL}/login`, userData)
}

export function getPost(postId: number) {
  return axios.get(`${API_URL}/posts/${postId}`)
}

export function getPosts() {
  return axios.get(`${API_URL}/posts`)
}

export function addView(postId: number) {
  return axios.put( `${API_URL}/posts/${postId}/views`, {})
}

export function addPost(postData, token: string) {
  return axios.post(`${API_URL}/posts`, postData, {
    headers: { Authorization: `Bearer: ${token}` },
  })
}

export function updatePost(postData, token: string) {
  return axios.put(`${API_URL}/posts/${postData.id}`, postData, {
    headers: { Authorization: `Bearer: ${token}` },
  })
}

export function getUserPosts(userId: number, token: string) {
  return axios.get(`${API_URL}/users/${userId}/posts`, {
    headers: { Authorization: `Bearer: ${token}` },
  })
}

export function updatePublishedStatus(
  postId: number,
  published: boolean,
  token: string
) {
  return axios.patch(
    `${API_URL}/posts/${postId}/published`,
    {
      published,
    },
    {
      headers: { Authorization: `Bearer: ${token}` },
    }
  )
}

export function deletePost(postId: number, token: string) {
  return axios.delete(`${API_URL}/posts/${postId}`, {
    headers: { Authorization: `Bearer: ${token}` },
  })
}
