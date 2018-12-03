pipeline is JUST one of the pattern in learning concurrecny and parallelism

each stage has ANY number of inbound and outbound channels, EXCEPT
for first and last
the first is only outbound
the last is only inbound

first produce 1/n channels
last consume 1/n channels
