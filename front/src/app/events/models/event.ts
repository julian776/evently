export type Event = {
  id: string;

  title: string;

  description: string;

  location: string;

  attendees?: string[]

  cost: number;

  organizerName: string;

  organizerEmail: string;

  startDate: Date | string

  endDate: Date | string

  startTime: string;

  endTime: string;
};
