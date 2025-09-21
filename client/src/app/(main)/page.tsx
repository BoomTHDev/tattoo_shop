import WorkCard from '@/components/work-card'

export default function HomePage() {
  const works = [
    {
      id: '1',
      title: 'Work 1',
      description: 'Description 1',
      imageUrl: 'https://picsum.photos/150',
    },
    {
      id: '2',
      title: 'Work 2',
      description: 'Description 2',
      imageUrl: 'https://picsum.photos/150',
    },
    {
      id: '3',
      title: 'Work 3',
      description: 'Description 3',
      imageUrl: 'https://picsum.photos/150',
    },
    {
      id: '4',
      title: 'Work 4',
      description: 'Description 4',
      imageUrl: 'https://picsum.photos/150',
    },
  ]

  return (
    <section className='grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8'>
      {works.map((work) => (
        <WorkCard key={work.id} work={work} />
      ))}
    </section>
  )
}
