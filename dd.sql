/**********************************************************
*
* Data Definition for Thread-Mail postgres store.
*
* Source: https://github.com/markmnl/tmail-store-postgres
*
***********************************************************/

create table msg(
    id            bigserial       primary key,
    pid           bigint          references msg (id),        
    ts            bigint          not null,
    from_addr     varchar(100)    not null,
    to_addr       varchar(100),
    topic         varchar(100)    not null,
    type          varchar(100)    not null,
    sha256        bytea           not null    unqiue,
    psha256       bytea                       unique,
    content       text            not null
);

create table msg_attachment(
    msg_id        bigint          references msg (id),
    filename      varchar(255)    not null,
    uri           varchar(1000)   not null,
    filesize      int             not null,
    token         varchar(1000),
    downloaded    timestamp,
    filepath      varchar(4096),
    primary key (msg_id, filename)
);