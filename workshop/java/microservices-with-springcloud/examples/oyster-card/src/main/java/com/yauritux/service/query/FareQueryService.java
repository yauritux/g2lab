package com.yauritux.service.query;

import java.util.List;
import java.util.Optional;

import com.yauritux.model.constant.JourneyType;
import com.yauritux.model.entity.Fare;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface FareQueryService {

	Optional<List<Fare>> findByJourneyType(JourneyType journeyType);
}
