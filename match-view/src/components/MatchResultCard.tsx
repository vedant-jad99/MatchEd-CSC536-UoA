
import { MatchResult } from "@/types/matching";
import {
  Card,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { CheckCircle, AlertTriangle, XCircle, ChevronDown, ChevronUp } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import MatchStatusBadge from "./MatchStatusBadge";
import { useState } from "react";
import { cn } from "@/lib/utils";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";

interface MatchResultCardProps {
  match: MatchResult;
}

const MatchResultCard = ({ match }: MatchResultCardProps) => {
  const [isOpen, setIsOpen] = useState(false);

  const getFormattedDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString();
  };

  const getFieldIcon = (match: boolean) => {
    if (match) {
      return <CheckCircle size={16} className="text-green-500" />;
    }
    return <XCircle size={16} className="text-red-500" />;
  };

  return (
    <Collapsible
      open={isOpen}
      onOpenChange={setIsOpen}
      className="w-full mb-2 rounded-md border"
    >
      <div className="flex items-center justify-between p-4">
        <div className="flex items-center gap-2 min-w-0">
          <div className="min-w-0 flex-1">
            <h4 className="font-medium text-sm truncate" title={match.course.name}>
              {match.course.name}
            </h4>
            <div className="flex items-center text-xs text-muted-foreground gap-1">
              <span className="truncate">{match.user.name}</span>
              <span className="mx-1">•</span>
              <span className="hidden sm:inline">ID: {match.course.ID}</span>
            </div>
          </div>
        </div>
        <div className="flex items-center gap-2">
          <MatchStatusBadge status={match.status} showIcon={false} className="text-xs py-0.5 h-5" />
          <CollapsibleTrigger asChild>
            <Button variant="ghost" size="sm" className="h-6 w-6 p-0">
              {isOpen ? <ChevronUp size={16} /> : <ChevronDown size={16} />}
            </Button>
          </CollapsibleTrigger>
        </div>
      </div>

      <CollapsibleContent>
        <Separator />
        <div className="px-4 py-2 bg-gray-50">
          <div className="flex items-center gap-1">
            <span className="text-xs font-medium">Match Confidence: {(match.confidence * 100).toFixed(0)}%</span>
            <div className="w-full bg-gray-200 rounded-full h-1.5">
              <div 
                className={cn(
                  "h-1.5 rounded-full",
                  match.status === 'perfect' ? "bg-matching-perfect" : 
                  match.status === 'partial' ? "bg-matching-partial" : 
                  "bg-matching-none"
                )}
                style={{ width: `${match.confidence * 100}%` }}
              ></div>
            </div>
          </div>
        </div>
        <CardContent className="pt-4 pb-2">
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <h5 className="text-xs font-semibold mb-2">Course Information</h5>
              <div className="space-y-1 text-sm">
                <div className="grid grid-cols-[80px_1fr] gap-2">
                  <span className="text-xs text-muted-foreground">Number:</span>
                  <span className="text-xs">{match.course.number}</span>
                </div>
                <div className="grid grid-cols-[80px_1fr] gap-2">
                  <span className="text-xs text-muted-foreground">Campus:</span>
                  <span className="text-xs">{match.course.campus}</span>
                </div>
                <div className="grid grid-cols-[80px_1fr] gap-2">
                  <span className="text-xs text-muted-foreground">Semesters:</span>
                  <span className="text-xs">{match.course.semesters.join(', ')}</span>
                </div>
              </div>
            </div>
            
            <div>
              <h5 className="text-xs font-semibold mb-2">User Information</h5>
              <div className="space-y-1 text-sm">
                <div className="grid grid-cols-[80px_1fr] gap-2">
                  <span className="text-xs text-muted-foreground">Name:</span>
                  <span className="text-xs">{match.user.name}</span>
                </div>
                <div className="grid grid-cols-[80px_1fr] gap-2">
                  <span className="text-xs text-muted-foreground">Email:</span>
                  <span className="text-xs">{match.user.email}</span>
                </div>
                <div className="grid grid-cols-[80px_1fr] gap-2">
                  <span className="text-xs text-muted-foreground">ID:</span>
                  <span className="text-xs">{match.user.ID}</span>
                </div>
              </div>
            </div>
          </div>

          {match.details && match.details.length > 0 && (
            <div className="mt-4">
              <h5 className="text-xs font-semibold mb-2">Match Details</h5>
              <div className="space-y-1">
                {match.details.map((detail, idx) => (
                  <div key={idx} className="grid grid-cols-[24px_1fr] gap-1">
                    <div className="flex items-center">
                      {getFieldIcon(detail.match)}
                    </div>
                    <div className="text-xs">
                      <span className="font-medium">{detail.field}: </span> 
                      {detail.match ? (
                        <span>{detail.sourceValue}</span>
                      ) : (
                        <div className="flex flex-col">
                          <span className="text-xs text-muted-foreground">Source: <span className="text-foreground">{detail.sourceValue}</span></span>
                          <span className="text-xs text-muted-foreground">Target: <span className="text-foreground">{detail.targetValue}</span></span>
                        </div>
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}
        </CardContent>
        
        <CardFooter className="py-2">
          <span className="text-xs text-gray-500">{getFormattedDate(match.timestamp)}</span>
        </CardFooter>
      </CollapsibleContent>
    </Collapsible>
  );
};

export default MatchResultCard;
