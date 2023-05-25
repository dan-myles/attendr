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
const TrackedClasses = async () => {
  const classes = await getTrackedClasses();

  if (!classes) return null;

  return (
    <div className="transition ease-in-out hover:-translate-y-1 hover:shadow-2xl">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Class</TableHead>
            <TableHead>Title</TableHead>
            <TableHead>Instructor</TableHead>
            <TableHead>Term</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {classes.map((c, i) => (
            <TableRow key={c.class_number}>
              <TableCell className="font-medium">
                {c.subject + " " + c.subject_number}
              </TableCell>
              <TableCell>{c.title}</TableCell>
              <TableCell>{c.instructor}</TableCell>
              {c.term === "2237" ? (
                <TableCell>Fall 2023</TableCell>
              ) : (
                <TableCell>{c.term}</TableCell>
              )}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};

export default TrackedClasses;
