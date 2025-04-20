
import * as React from "react";
import { cn } from "@/lib/utils";
import {
  ChevronLeft,
  ChevronRight,
  ChevronsLeft,
  ChevronsRight,
  MoreHorizontal,
} from "lucide-react";
import { Button } from "@/components/ui/button";

const Pagination = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    role="navigation"
    aria-label="pagination"
    className={cn("flex justify-center", className)}
    {...props}
  />
);

const PaginationItem = ({
  className,
  active,
  ...props
}: React.ComponentProps<typeof Button> & {
  active?: boolean;
}) => (
  <Button
    aria-current={active ? "page" : undefined}
    variant={active ? "default" : "outline"}
    size="icon"
    className={cn(
      "h-8 w-8 text-xs",
      {
        "pointer-events-none": active,
      },
      className
    )}
    {...props}
  />
);

const PaginationEllipsis = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLSpanElement>) => (
  <span
    aria-hidden
    className={cn("flex h-8 w-8 items-center justify-center", className)}
    {...props}
  >
    <MoreHorizontal className="h-4 w-4" />
    <span className="sr-only">More pages</span>
  </span>
);

const PaginationPrev = ({
  className,
  ...props
}: React.ComponentProps<typeof Button>) => (
  <Button
    aria-label="Go to previous page"
    size="icon"
    variant="outline"
    className={cn("h-8 w-8", className)}
    {...props}
  >
    <ChevronLeft className="h-4 w-4" />
  </Button>
);

const PaginationNext = ({
  className,
  ...props
}: React.ComponentProps<typeof Button>) => (
  <Button
    aria-label="Go to next page"
    size="icon"
    variant="outline"
    className={cn("h-8 w-8", className)}
    {...props}
  >
    <ChevronRight className="h-4 w-4" />
  </Button>
);

const PaginationFirst = ({
  className,
  ...props
}: React.ComponentProps<typeof Button>) => (
  <Button
    aria-label="Go to first page"
    size="icon"
    variant="outline"
    className={cn("h-8 w-8", className)}
    {...props}
  >
    <ChevronsLeft className="h-4 w-4" />
  </Button>
);

const PaginationLast = ({
  className,
  ...props
}: React.ComponentProps<typeof Button>) => (
  <Button
    aria-label="Go to last page"
    size="icon"
    variant="outline"
    className={cn("h-8 w-8", className)}
    {...props}
  >
    <ChevronsRight className="h-4 w-4" />
  </Button>
);

Pagination.Item = PaginationItem;
Pagination.Ellipsis = PaginationEllipsis;
Pagination.Prev = PaginationPrev;
Pagination.Next = PaginationNext;
Pagination.First = PaginationFirst;
Pagination.Last = PaginationLast;

export { Pagination };
