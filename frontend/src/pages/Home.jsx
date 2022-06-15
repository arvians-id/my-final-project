import React from 'react'
import useStore from '../provider/zustand/store';

export default function Home() {
  const user = useStore((state) => state.user);

  return (
    <div>
        <h1>Hy {user.username}, Selamat Datang Di Halaman Home</h1>
    </div>
  )
}
