import { auth } from "@clerk/nextjs";
import { PrismaClient } from "@prisma/client";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

const TrackedClasses = async () => {
  async function getTrackedClasses() {
    "use server";

    const { userId } = auth();
    if (!userId) {
      throw new Error("You must be signed in to access tracked classes!");
    }

    const prisma = new PrismaClient();
    const trackedClasses = await prisma.aSU_Watched_Classes.findMany({
      where: {
        user_id: userId,
      },
    });

    return trackedClasses;
  }

  const classes = await getTrackedClasses();

  if (!classes) return null;

  return (
    <div>
      <p>LENGTH={classes.length}</p>
      <Table>
        <TableCaption>A list of your recent invoices.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Class</TableHead>
            <TableHead>Title</TableHead>
            <TableHead>Instructor</TableHead>
            <TableHead>Term</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow>
            {classes.map((c, i) => (
              <TableRow key={c.user_id}>
                <TableCell className="font-medium">
                  {c.subject + c.subject_number}
                </TableCell>
                <TableCell>{c.title}</TableCell>
                <TableCell>{c.instructor}</TableCell>
                <TableCell>{c.term}</TableCell>
              </TableRow>
            ))}
          </TableRow>
        </TableBody>
      </Table>
    </div>
  );
};

export default TrackedClasses;
