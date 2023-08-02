# DHCP Address Squatter
A gross hack.

## WHY?!?!?!?!?!?!?!?!
The TELUS Wi-Fi Hub router provided by my ISP does not allow me to set a DNS server
address that is within the local network to be sent in DHCP Offers. This is undocumented
and if one tries to do so the router will just save the setting but ignore it and send
whatever it likes (thanks Telus, for all the hours I spent wondering why the hell my
split DNS didn't work).

So, solution is: disable DHCP on the ISP router and move it to the local DNS server
that I'm already using. The problem is, the router doesn't allow you to disable DHCP
if it isn't in bridge mode. What I can do is restrict the DHCP assignment range to a
single IP.

Now I have most DHCP assignments coming from my preferred server, but occasionally one
device gets an offer from the ISP router and happily takes it, creating a weird partition.

That's where this lil guy comes in, it constantly sends out DHCP requests until it gets the
particular 'bad' offer and squats on it immediately forever :)

## Why can't you just....get a better router?
That costs money. Also I'm a bit bored.



