package com.yauritux.service.query;

import java.util.Optional;

import com.yauritux.model.entity.Station;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface StationQueryService {

	Optional<Station> findByName(final String name);
}
