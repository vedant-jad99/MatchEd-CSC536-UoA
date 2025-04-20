
import { MatchStatus } from "@/types/matching";
import { cn } from "@/lib/utils";
import { Badge } from "@/components/ui/badge";
import { CheckCircle2, AlertCircle, XCircle } from "lucide-react";

interface MatchStatusBadgeProps {
  status: MatchStatus;
  showIcon?: boolean;
  className?: string;
}

const MatchStatusBadge = ({ status, showIcon = true, className }: MatchStatusBadgeProps) => {
  const config = {
    perfect: {
      label: "Perfect Match",
      color: "bg-matching-perfect text-white hover:bg-matching-perfect/90",
      icon: CheckCircle2
    },
    partial: {
      label: "Partial Match",
      color: "bg-matching-partial text-white hover:bg-matching-partial/90",
      icon: AlertCircle
    },
    none: {
      label: "No Match",
      color: "bg-matching-none text-white hover:bg-matching-none/90",
      icon: XCircle
    }
  };

  const { label, color, icon: Icon } = config[status];

  return (
    <Badge variant="outline" className={cn("font-medium", color, className)}>
      {showIcon && <Icon className="mr-1 h-3 w-3" />}
      {label}
    </Badge>
  );
};

export default MatchStatusBadge;
