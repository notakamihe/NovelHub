import moment from "moment"

export const formatISBN = (str) => {
    if (str && str.length == 13) {
        const segments = [str.slice(0, 3), str[3], str.slice(4, 8), str.slice(8, 12), str[12]]
        return segments.join("-")
    }

    return str
}

export const getFileName = (url) => {
    if (!url)
        return

    const split = url.split(/[\\/]/)
    return split[split.length - 1]
}

export const normalizeUrl = (url) => {
    return [url.slice(0, url.indexOf("\\")), "\\", url.slice(url.indexOf("\\"))].join("")
}

export const timeAgo = (date) => {
    moment.updateLocale("en", {
        relativeTime: {
            future: '%s',
            past: '%s',
            s: '%ds',
            ss: '%ds',
            m: '%dm',
            mm: '%dm',
            w: "%dw",
            ww: "%dw",
            h: '%dh',
            hh: '%dh',
            d: '%dd',
            dd: '%dd',
            y: "%dy",
            yy: "%dy"
        }
    })

    return moment(date).fromNow()
}