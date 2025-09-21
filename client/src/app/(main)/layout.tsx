import Navbar from '@/components/layout/navbar'

const MainLayout = (props: LayoutProps<'/'>) => {
  return (
    <main className='container mx-auto px-4 py-12'>
      <Navbar />
      {props.children}
    </main>
  )
}
export default MainLayout
