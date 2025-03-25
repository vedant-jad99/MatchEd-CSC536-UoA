
import React from 'react';
import { MainLayout } from '@/components/layout/MainLayout';
import { 
  Card, 
  CardContent, 
  CardDescription, 
  CardFooter, 
  CardHeader, 
  CardTitle 
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Separator } from '@/components/ui/separator';
import { User, Mail, Settings, ShieldCheck, Bell } from 'lucide-react';
import { Switch } from '@/components/ui/switch';

const Profile = () => {
  return (
    <MainLayout>
      <div className="space-y-6">
        <div>
          <h1 className="text-2xl font-bold tracking-tight">Profile & Settings</h1>
          <p className="text-muted-foreground">Manage your account and preferences</p>
        </div>
        
        <Tabs defaultValue="profile" className="space-y-6">
          <TabsList className="mb-2">
            <TabsTrigger value="profile">
              <User className="h-4 w-4 mr-2" />
              Profile
            </TabsTrigger>
            <TabsTrigger value="account">
              <Settings className="h-4 w-4 mr-2" />
              Account
            </TabsTrigger>
            <TabsTrigger value="notifications">
              <Bell className="h-4 w-4 mr-2" />
              Notifications
            </TabsTrigger>
          </TabsList>
          
          <TabsContent value="profile">
            <div className="grid gap-6 md:grid-cols-2">
              <Card className="animate-fade-in">
                <CardHeader>
                  <CardTitle>Personal Information</CardTitle>
                  <CardDescription>Update your personal details here</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="flex flex-col md:flex-row md:items-center gap-4">
                    <Avatar className="h-20 w-20 border">
                      <AvatarFallback className="text-lg">JD</AvatarFallback>
                    </Avatar>
                    <div className="space-y-1 flex-1">
                      <h4 className="text-sm font-medium">Profile Photo</h4>
                      <p className="text-sm text-muted-foreground">Upload a new profile picture</p>
                      <div className="flex items-center gap-2 mt-2">
                        <Button size="sm">Upload</Button>
                        <Button size="sm" variant="outline">Remove</Button>
                      </div>
                    </div>
                  </div>
                  
                  <Separator />
                  
                  <div className="space-y-4">
                    <div className="grid gap-2">
                      <Label htmlFor="name">Full Name</Label>
                      <Input id="name" defaultValue="John Doe" />
                    </div>
                    <div className="grid gap-2">
                      <Label htmlFor="email">Email</Label>
                      <Input id="email" type="email" defaultValue="john@example.com" />
                    </div>
                    <div className="grid gap-2">
                      <Label htmlFor="role">Job Title</Label>
                      <Input id="role" defaultValue="Software Engineer" />
                    </div>
                  </div>
                </CardContent>
                <CardFooter className="flex justify-end gap-2">
                  <Button variant="outline">Cancel</Button>
                  <Button>Save Changes</Button>
                </CardFooter>
              </Card>
              
              <Card className="animate-fade-in" style={{animationDelay: '100ms'}}>
                <CardHeader>
                  <CardTitle>Work Information</CardTitle>
                  <CardDescription>Update your professional details</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="grid gap-2">
                    <Label htmlFor="company">Company</Label>
                    <Input id="company" defaultValue="Acme Inc." />
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="department">Department</Label>
                    <Input id="department" defaultValue="Engineering" />
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="skills">Skills (comma separated)</Label>
                    <Input id="skills" defaultValue="JavaScript, React, UI Design" />
                  </div>
                </CardContent>
                <CardFooter className="flex justify-end gap-2">
                  <Button variant="outline">Cancel</Button>
                  <Button>Save Changes</Button>
                </CardFooter>
              </Card>
            </div>
          </TabsContent>
          
          <TabsContent value="account">
            <div className="grid gap-6 md:grid-cols-2">
              <Card className="animate-fade-in">
                <CardHeader>
                  <CardTitle>Password</CardTitle>
                  <CardDescription>Update your password</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="grid gap-2">
                    <Label htmlFor="current-password">Current Password</Label>
                    <Input id="current-password" type="password" />
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="new-password">New Password</Label>
                    <Input id="new-password" type="password" />
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="confirm-password">Confirm New Password</Label>
                    <Input id="confirm-password" type="password" />
                  </div>
                </CardContent>
                <CardFooter className="flex justify-end gap-2">
                  <Button variant="outline">Cancel</Button>
                  <Button>Update Password</Button>
                </CardFooter>
              </Card>
              
              <Card className="animate-fade-in" style={{animationDelay: '100ms'}}>
                <CardHeader>
                  <CardTitle>Security</CardTitle>
                  <CardDescription>Manage your security settings</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="flex items-center justify-between">
                    <div className="space-y-0.5">
                      <Label>Two-Factor Authentication</Label>
                      <p className="text-sm text-muted-foreground">Add an extra layer of security</p>
                    </div>
                    <Switch />
                  </div>
                  <Separator />
                  <div className="flex items-center justify-between">
                    <div className="space-y-0.5">
                      <Label>Session Management</Label>
                      <p className="text-sm text-muted-foreground">Manage active sessions</p>
                    </div>
                    <Button variant="outline" size="sm">View Sessions</Button>
                  </div>
                </CardContent>
                <CardFooter>
                  <div className="flex items-center space-x-2 text-sm text-muted-foreground">
                    <ShieldCheck className="h-4 w-4" />
                    <p>Last security audit: 2 weeks ago</p>
                  </div>
                </CardFooter>
              </Card>
            </div>
          </TabsContent>
          
          <TabsContent value="notifications">
            <Card className="animate-fade-in">
              <CardHeader>
                <CardTitle>Notification Preferences</CardTitle>
                <CardDescription>Manage how you receive notifications</CardDescription>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="space-y-4">
                  <h3 className="text-sm font-medium">Email Notifications</h3>
                  <div className="grid gap-2">
                    <div className="flex items-center justify-between">
                      <Label htmlFor="email-roles" className="flex items-center gap-2 cursor-pointer">
                        Role Assignments
                        <p className="text-sm font-normal text-muted-foreground">Receive notifications when roles are assigned or changed</p>
                      </Label>
                      <Switch id="email-roles" defaultChecked />
                    </div>
                    <Separator />
                    <div className="flex items-center justify-between">
                      <Label htmlFor="email-updates" className="flex items-center gap-2 cursor-pointer">
                        Product Updates
                        <p className="text-sm font-normal text-muted-foreground">Receive notifications about new features and updates</p>
                      </Label>
                      <Switch id="email-updates" />
                    </div>
                    <Separator />
                    <div className="flex items-center justify-between">
                      <Label htmlFor="email-security" className="flex items-center gap-2 cursor-pointer">
                        Security Alerts
                        <p className="text-sm font-normal text-muted-foreground">Receive notifications about security-related events</p>
                      </Label>
                      <Switch id="email-security" defaultChecked />
                    </div>
                  </div>
                </div>
                
                <div className="space-y-4">
                  <h3 className="text-sm font-medium">In-App Notifications</h3>
                  <div className="grid gap-2">
                    <div className="flex items-center justify-between">
                      <Label htmlFor="inapp-roles" className="flex items-center gap-2 cursor-pointer">
                        Role Assignments
                        <p className="text-sm font-normal text-muted-foreground">Show notifications when roles are assigned or changed</p>
                      </Label>
                      <Switch id="inapp-roles" defaultChecked />
                    </div>
                    <Separator />
                    <div className="flex items-center justify-between">
                      <Label htmlFor="inapp-updates" className="flex items-center gap-2 cursor-pointer">
                        Product Updates
                        <p className="text-sm font-normal text-muted-foreground">Show notifications about new features and updates</p>
                      </Label>
                      <Switch id="inapp-updates" defaultChecked />
                    </div>
                    <Separator />
                    <div className="flex items-center justify-between">
                      <Label htmlFor="inapp-security" className="flex items-center gap-2 cursor-pointer">
                        Security Alerts
                        <p className="text-sm font-normal text-muted-foreground">Show notifications about security-related events</p>
                      </Label>
                      <Switch id="inapp-security" defaultChecked />
                    </div>
                  </div>
                </div>
              </CardContent>
              <CardFooter className="flex justify-end gap-2">
                <Button variant="outline">Reset to Defaults</Button>
                <Button>Save Preferences</Button>
              </CardFooter>
            </Card>
          </TabsContent>
        </Tabs>
      </div>
    </MainLayout>
  );
};

export default Profile;
