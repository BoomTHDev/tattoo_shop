import Link from 'next/link'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import Image from 'next/image'

interface WorkCardProps {
  work: Work
}

export default function WorkCard({ work }: WorkCardProps) {
  return (
    <Link href={`/reviews/${work.id}`}>
      <Card className='h-full overflow-hidden transition-transform duration-300 ease-in-out hover:scale-105 hover:shadow-lg'>
        <CardHeader>
          <div className='relative h-48 w-full'>
            <Image
              alt={work.title}
              src={work.imageUrl}
              fill
              className='object-cover rounded-t-lg'
              priority
            />
          </div>
        </CardHeader>
        <CardContent>
          <CardTitle className='mb-2 text-xl'>{work.title}</CardTitle>
          <CardDescription>{work.description}</CardDescription>
        </CardContent>
      </Card>
    </Link>
  )
}
